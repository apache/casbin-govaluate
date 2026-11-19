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

func TestEmptyInClause(t *testing.T) {
	for _, input := range []string{"false in ()", "'hello' in ( )", "value in ()"} {
		t.Run(input, func(t *testing.T) {
			expression, err := NewEvaluableExpression(input)
			if err != nil {
				t.Fatal(err)
			}
			_, err = expression.Evaluate(map[string]interface{}{"value": "hello"})
			if err == nil {
				t.Fatal("expected a type error for an empty clause")
			}
		})
	}
}

func TestInClauseAndEmptyFunctionCompatibility(t *testing.T) {
	functions := map[string]ExpressionFunction{
		"one": func(args ...interface{}) (interface{}, error) {
			if len(args) != 0 {
				t.Fatalf("got %d arguments; want zero", len(args))
			}
			return float64(1), nil
		},
	}
	for _, input := range []string{"1 in (1)", "1 in (1, 2)", "one() == 1"} {
		expression, err := NewEvaluableExpressionWithFunctions(input, functions)
		if err != nil {
			t.Fatal(err)
		}
		got, err := expression.Evaluate(nil)
		if err != nil || got != true {
			t.Fatalf("%s: got %v, error %v; want true", input, got, err)
		}
	}
}
