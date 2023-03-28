// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package precompile

// import (
// 	"fmt"
// 	"math/big"

// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// 	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
// )

// var _ = Describe("Distribution", func() {
// 	It("should call functions on the precompile directly", func() {
// 		tx, err := distributionPrecompile.GetWithdrawEnabled(nil)
// 		Expect(err).ToNot(HaveOccurred())
// 		Expect(tx).To(BeTrue())

// 		fmt.Println("First Call")

// 		txr := tf.GenerateTransactOpts("")
// 		txr.GasLimit = 10e6
// 		txr.Value = big.NewInt(10)
// 		a, err := distributionPrecompile.SetWithdrawAddress(txr, validator)
// 		Expect(err).ToNot(HaveOccurred())
// 		Expect(a).ToNot(BeNil())
// 		ExpectMined(tf.EthClient, a)
// 		ExpectSuccessReceipt(tf.EthClient, a)
// 	})
// })
