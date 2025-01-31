#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// The  functions defining the operations of the ALU
int add(int op1, int op2) { return op1 + op2; } // Adds two integers
int sub(int op1, int op2) { return op1 - op2; } // Subtracts the second integer from the first
int Div(int op1, int op2) { return op1 / op2; } // Divides the first integer by the second
int mult(int op1, int op2) { return op1 * op2; } // Multiplies two integers
int Xor(int op1, int op2) { return op1 ^ op2; } // Performs a bitwise XOR operation
int And(int op1, int op2) { return op1 & op2; } // Performs a bitwise AND operation
int Or(int op1, int op2) { return op1 | op2; } // Performs a bitwise OR operation

// The function which performs the alu operations
int ALU(const char *instruction) {
  int operand1, operand2;
  char operation;
  sscanf(instruction, "%c %d %d", &operation, &operand1, &operand2); // Reads the operation and operands from the instruction string
  switch (operation) {
  case 'a':
    return add(operand1, operand2); // Calls the add function
    break;

  case 's':
    return sub(operand1, operand2); // Calls the subtract function
    break;

  case 'd':
    return Div(operand1, operand2); // Calls the divide function
    break;

  case 'm':
    return mult(operand1, operand2); // Calls the multiply function
    break;
  case 'A':
    return And(operand1, operand2); // Calls the bitwise AND function
    break;

  case 'O':
    return Or(operand1, operand2); // Calls the bitwise OR function
    break;

  case 'X':
    return Xor(operand1, operand2); // Calls the bitwise XOR function
    break;
  default:
    return 0; // Returns 0 if the operation is not recognized
    break;
  }
}

// prints on how to use the alu
void print_help() {
  printf("-----------Help-----------\n");

  printf("The alu support the following operations: \n");
  printf("add: a\n");
  printf("mult: m\n");
  printf("sub: s\n");
  printf("div: d\n");
  printf("bitwise and: A\n");
  printf("bitwise or: O\n");
  printf("bitwise xor: X\n");

  printf("------------------------------------------------\n");
}

int main(void) {
  // 20 bytes to hold the instruction
  char instruction[20];
  // infinite loop
  while (1) {
    printf("Enter instruction in the format (eg a 10 5) or 'help' for more "
           "info or 'exit' or  to "
           "quit\n ");
    printf(">> ");
    fgets(instruction, 20, stdin); // Reads the instruction from the user
    if (strncmp(instruction, "exit", 4) == 0) { // Checks if the user wants to exit
      break;
    }
    if (strncmp(instruction, "help", 4) == 0) { // Checks if the user needs help
      print_help(); // Prints the help message
      continue;
    }
    int result = ALU(instruction); // Calls the ALU function with the instruction
    printf("------------------------------------------------\n");
    printf("The result is: %d\n", result); // Prints the result
    printf("------------------------------------------------\n");
  }

  return EXIT_SUCCESS;
}