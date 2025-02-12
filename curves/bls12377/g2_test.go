// Copyright 2023 Ingonyama
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bls12377

import (
	"fmt"
	"testing"

	icicle "github.com/ingonyama-zk/icicle/goicicle/curves/bls12377"
	"github.com/stretchr/testify/assert"
)

func TestToGnarkJacG2(t *testing.T) {
	gnark, _ := randG2Jac()

	var pointAffine icicle.G2PointAffine
	G2PointAffineFromGnarkJac(&gnark, &pointAffine)
	pointProjective := pointAffine.ToProjective()
	fmt.Printf("%+v\n", pointProjective)
	backToGnark := G2PointToGnarkJac(&pointProjective)

	assert.True(t, gnark.Equal(backToGnark))
}
