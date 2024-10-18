#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

typedef struct stackNode {
    char val;
    struct stackNode *next;
} stackNode; 

void push(stackNode **head, char value){
    stackNode *newNode = (stackNode*)malloc(sizeof(stackNode));
    newNode->val = value;
    newNode->next = *head;

    *head = newNode;
}

char pop(stackNode **head){
    if (*head == NULL)
    {
        return '\0';
    }
    
    stackNode *temp = *head;
    char poppedItem = temp->val;

    *head = (*head)->next;
    free(temp);

    return poppedItem;
}

bool isEmpty(stackNode *head){
    return head == NULL;
}

bool isValid(char *s){
    if (strlen(s) == 0)
    {
        return true;
    }

    stackNode *stack = NULL;
    bool valid = true;
    char lastItem;

    for (short i = 0; i < strlen(s); i++)
    {
        switch (s[i])
        {
        case '(':
            push(&stack, '(');
            break;
        case '[':
            push(&stack, '[');
            break;
        case '{':
            push(&stack, '{');
            break;
        case ')':
            lastItem = pop(&stack);
            if (lastItem != '(')
            {
                valid = false;
            }
            break;
        case ']':
            lastItem = pop(&stack);
            if (lastItem != '[')
            {
                valid = false;
            }
            break;
        case '}':
            lastItem = pop(&stack);
            if (lastItem != '{')
            {
                valid = false;
            }
            break;
        }
    }

    if (!isEmpty(stack))
    {
        valid = false;
    }

    return valid;
}

int main(){
    struct stackNode *stack = NULL;  // Initial ist der Stack leer

    // Werte zum Stack hinzufügen
    push(&stack, '(');
    push(&stack, '{');

    // Wert entfernen (pop)
    char topValue = pop(&stack);
    printf("Popped value: %c\n", topValue);
}