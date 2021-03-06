package preprocessor

import (
	"fmt"
	"testing"
)

func TestParseIncludeList(t *testing.T) {
	testCases := []struct {
		inputLine string
		list      []string
	}{
		{
			inputLine: ` exit.o: exit.c tests.h `,
			list:      []string{"exit.c", "tests.h"},
		},
		{

			inputLine: ` exit.o: exit.c /usr/include/stdlib.h /usr/include/features.h \
  /usr/include/stdc-predef.h /usr/include/x86_64-linux-gnu/sys/cdefs.h \
  /usr/include/x86_64-linux-gnu/gnu/stubs-64.h \
  /usr/lib/llvm-3.8/bin/../lib/clang/3.8.0/include/stddef.h
  `,
			list: []string{"exit.c", "/usr/include/stdlib.h", "/usr/include/features.h",
				"/usr/include/stdc-predef.h", "/usr/include/x86_64-linux-gnu/sys/cdefs.h",
				"/usr/include/x86_64-linux-gnu/gnu/stubs-64.h",
				"/usr/lib/llvm-3.8/bin/../lib/clang/3.8.0/include/stddef.h",
			},
		},
		{
			inputLine: ` main.o: \
  /home/Глава\ 6/6\ .2/main.c  /home/e\ 1.c`,
			list: []string{"/home/Глава 6/6 .2/main.c", "/home/e 1.c"},
		},
		{
			inputLine: ` main.o: \
  /home/lepricon/go/src/github.com/FTwOoO/c4go/build/git-source/VasielBook/Глава\ 6/6.2/main.c`,
			list: []string{"/home/lepricon/go/src/github.com/FTwOoO/c4go/build/git-source/VasielBook/Глава 6/6.2/main.c"},
		},
		{
			inputLine: ` shell.o: /tmp/SQLITE/sqlite-amalgamation-3220000/shell.c \
  /tmp/SQLITE/sqlite-amalgamation-3220000/sqlite3.h
sqlite3.o: /tmp/SQLITE/sqlite-amalgamation-3220000/sqlite3.c /tmp/1.h \
 /tmp/2.h`,
			list: []string{
				"/tmp/SQLITE/sqlite-amalgamation-3220000/shell.c",
				"/tmp/SQLITE/sqlite-amalgamation-3220000/sqlite3.h",
				"/tmp/SQLITE/sqlite-amalgamation-3220000/sqlite3.c",
				"/tmp/1.h",
				"/tmp/2.h",
			},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test:%d", i), func(t *testing.T) {
			actual, err := parseIncludeList(tc.inputLine)
			if err != nil {
				t.Fatal(err)
			}
			if len(actual) != len(tc.list) {
				t.Fatalf("Cannot parse line : %s.\nActual result : %#v.\nExpected: %#v", tc.inputLine, actual, tc.list)
			}
			for i := range actual {
				if actual[i] != tc.list[i] {
					t.Fatalf("Cannot parse 'include' in line : %s.\nActual result : %#v.\nExpected: %#v", tc.inputLine, actual[i], tc.list[i])
				}
			}
		})
	}
}

func TestIncludeListFail(t *testing.T) {
	_, err := parseIncludeList("without index")
	if err == nil {
		t.Fatalf("Cannot error withou :")
	}
}
