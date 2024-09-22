//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package debug

// Field is a helper function for creating a Node with a label and one child label.
// This is how a field with a value is typically represented.
func Field(name string, value string) Node {
	return Node{
		Label: name,
		Children: Nodes{
			Node{
				Label: value,
			},
		},
	}
}

// Nested is a helper function for creating a Node.
func Nested(name string, fs ...Node) Node {
	return Node{
		Label:    name,
		Children: Nodes(fs),
	}
}
