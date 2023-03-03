#include <stdio.h>
#include <unistd.h>

int main() {
	printf("There should be nothing printed after this!\n");

	// noop is an empty executable file
	char* a[] = {"./noop", NULL};
	execvp(a[0], a);

	printf("This line should not be printed after noop!\n");
}

