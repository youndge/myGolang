#include "344.h"
void swap(char *a, char *b) {
	char t = *a;
	*a = *b, *b = t;
}

void reverseString(char *s, int sSize) {
	for (int left = 0, right = sSize - 1; left < right; ++left, --right) {
		swap(s + left, s + right);
	}
}

void test344() {
	char s[5] = { 'H','e','l','l','o' };
	for (int i = 0; i < 5; i++) {
		printf("%c  ",s[i]);
	}
	printf_s("\n");
	reverseString(s, 5);
	for (int i = 0; i < 5; i++) {
		printf_s("%c  ", s[i]);
	}
	printf_s("\n");
}
