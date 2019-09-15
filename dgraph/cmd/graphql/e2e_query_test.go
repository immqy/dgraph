/*
 * Copyright 2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package graphql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func getCountry(t *testing.T, countryUID, expected string) {
	getCountryParams := &GraphQLParams{
		Query: `query getCountry($createdID: ID!) {
			getCountry(id: $createdID) {
				id
				name
			}
		}`,
		Variables: map[string]interface{}{"createdID": countryUID},
	}

	resp, err := getCountryParams.ExecuteAsPost(graphqlURL)
	require.NoError(t, err)

	require.JSONEq(t, expected, string(resp))
}

func queryCountryByRegExp(t *testing.T, regexp, expected string) {
	getCountryParams := &GraphQLParams{
		Query: `query queryCountry($regexp: String!) {
			queryCountry(filter: { name: { regexp: $regexp } }) {
				name
			}
		}`,
		Variables: map[string]interface{}{"regexp": regexp},
	}

	resp, err := getCountryParams.ExecuteAsPost(graphqlURL)
	require.NoError(t, err)

	require.JSONEq(t, expected, string(resp))
}
