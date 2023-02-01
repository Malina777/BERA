// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package snapshot

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	"github.com/berachain/stargazer/lib/registry"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `initJournalCapacity` is the initial capacity of the `journal` stack.
// TODO: determine better initial capacity.
const initJournalCapacity = 32

// `revision` is a snapshot revision, which holds all `Controllable`s' snapshot ids.
// Specifically, it is a map of a `Controllable`'s `RegistryKey` to its corresponding current
// snapshot revision id.
type revision[K comparable] map[K]int

// `controller` conforms to the `libtypes.Controller` interface and is used to register and sync
// snapshotting across multiple `libtypes.Controllable` objects.
type controller[K comparable, T libtypes.Controllable[K]] struct {
	// `Registry` stores the `Controllable` objects.
	libtypes.Registry[K, T]

	// `journal` is a stack of `revision`s.
	journal ds.Stack[revision[K]]
}

// `NewController` returns a new `Controller` object.
func NewController[K comparable, T libtypes.Controllable[K]]() libtypes.Controller[K, T] {
	return &controller[K, T]{
		Registry: registry.NewMap[K, T](),
		journal:  stack.New[revision[K]](initJournalCapacity),
	}
}

// `Snapshot` takes a snapshot for all controllable objects and returns the controller's snapshot
// id.
//
// `Snapshot` implements `libtypes.Snapshottable`.
func (c *controller[K, T]) Snapshot() int {
	snap := make(revision[K])
	for key, controllable := range c.Iterate() {
		snap[key] = controllable.Snapshot()
	}
	c.journal.Push(snap)

	return c.journal.Size()
}

// `RevertToSnapshot` reverts all controllable objects to their own snapshot id corresponding to
// `id`.
//
// `RevertToSnapshot` implements `libtypes.Snapshottable`.
func (c *controller[K, T]) RevertToSnapshot(id int) {
	lastestSnapshot := c.journal.Peek()
	for key, controllable := range c.Iterate() {
		// Only revert if exists. This is to handle the case where a `libtypes.Controllable` object
		// is added after a snapshot has been taken.
		if revision, ok := lastestSnapshot[key]; ok {
			controllable.RevertToSnapshot(revision)
		}
	}
	c.journal.PopToSize(id)
}

// `Finalize` writes all the controllables controlled by this controller.
//
// `Finalize` implements `libtypes.Controller`.
func (c *controller[K, T]) Finalize() {
	for _, controllable := range c.Iterate() {
		controllable.Write()
	}
}
