package isPalindrome

import "testing"

//回文数
func TestIsPalindrome(t *testing.T) {
	if r := IsPalindrome(12321); !r {
		t.Errorf("12321 Is the number of palindrome，but %t", r)
	}
	if r := IsPalindrome(1221); !r {
		t.Errorf("1221 Is the number of palindrome，but %t", r)
	}
	if r := IsPalindrome(121); !r {
		t.Errorf("121 Is the number of palindrome，but %t", r)
	}
	if r := IsPalindrome(-121); r {
		t.Errorf("-121 Is not the number of palindrome，but %t", r)
	}
	if r := IsPalindrome(100); r {
		t.Errorf("100 Is not the number of palindrome，but %t", r)
	}
}

func TestLeetcodeIsPalindrome(t *testing.T) {
	if r := LeetcodeIsPalindrome(12321); !r {
		t.Errorf("12321 Is the number of palindrome，but %t", r)
	}
	if r := LeetcodeIsPalindrome(1221); !r {
		t.Errorf("1221 Is the number of palindrome，but %t", r)
	}
	if r := LeetcodeIsPalindrome(121); !r {
		t.Errorf("121 Is the number of palindrome，but %t", r)
	}
	if r := LeetcodeIsPalindrome(-121); r {
		t.Errorf("-121 Is not the number of palindrome，but %t", r)
	}
	if r := LeetcodeIsPalindrome(100); r {
		t.Errorf("100 Is not the number of palindrome，but %t", r)
	}
}
