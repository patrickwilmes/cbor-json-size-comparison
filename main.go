/*
 * SPDX-License-Identifier: CC-BY-NC-ND-4.0
 *
 * cbor-json-size-comparison
 * Copyright (c) 2024 Patrick Wilmes <p.wilmes89@gmail.com>
 *
 * This file is licensed under the Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International License.
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://creativecommons.org/licenses/by-nc-nd/4.0/
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fxamacker/cbor/v2"
)

type Level1 struct {
	Description string `json:",omitempty"`
}

type Level2 struct {
	PreviousLevel Level1 `json:",omitempty"`
	Description   string `json:",omitempty"`
}

type Level3 struct {
	PreviousLevel Level2 `json:",omitempty"`
	Description   string `json:",omitempty"`
}

func encodeUsingCBOR(level3 Level3) {
	results, _ := cbor.Marshal(level3)
	fmt.Println("hex(CBOR): " + hex.EncodeToString(results))
	text, _ := cbor.Diagnose(results)
	fmt.Println("text: " + text)
	fmt.Printf("CBOR SIZE IN BYTES: %d\n", len(results))
}

func encodeUsingJson(level3 Level3) {
	results, _ := json.Marshal(level3)
	fmt.Println("json: " + hex.EncodeToString(results))

	text := string(results)
	fmt.Println("text: " + text)
	fmt.Printf("JSON SIZE IN BYTES: %d\n", len(results))
}

func main() {
	structure := Level3{
		PreviousLevel: Level2{
			PreviousLevel: Level1{Description: "Another Hello World"},
			Description:   "And Hello World again",
		},
		Description: "HelloWorld",
	}
	encodeUsingCBOR(structure)
	fmt.Println("-----------------------------------------------")
	encodeUsingJson(structure)
}
