// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied. See the License for the
// specific language governing permissions and limitations
// under the License.

package govaluate

import "testing"

func TestBitwisePrecedenceWithParameters(t *testing.T) {
	cases := []struct {
		expression string
		want       interface{}
	}{
		{"a | b & c", float64(1)},
		{"(a | b) & c", float64(0)},
		{"a | (b & c)", float64(1)},
		{"a | a ^ a", float64(1)},
		{"(a | a) ^ a", float64(0)},
		{"c ^ b & a", float64(4)},
		{"(c ^ b) & a", float64(0)},
		{"c & (c | b) ^ a", float64(5)},
		{"a | b << a", float64(5)},
		{"(a | b) << a", float64(6)},
		{"a | b + c", float64(7)},
		{"(a | b) + c", float64(7)},
		{"(a | b & c) == a", true},
		{"((a | b) & c) == a", false},
	}
	for _, tc := range cases {
		t.Run(tc.expression, func(t *testing.T) {
			expression, err := NewEvaluableExpression(tc.expression)
			if err != nil {
				t.Fatal(err)
			}
			// Evaluate the same parsed expression repeatedly, including cache builds.
			for i := 0; i < 3; i++ {
				got, err := expression.Evaluate(map[string]interface{}{
					"a": float64(1), "b": float64(2), "c": float64(4),
				})
				if err != nil || got != tc.want {
					t.Fatalf("got %v, error %v; want %v", got, err, tc.want)
				}
			}
		})
	}
}
