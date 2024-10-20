#include <stdio.h>
#include <stdlib.h>

typedef struct ListNode ListNode; // Vorwärtsdeklaration

struct ListNode
{
    int val;
    ListNode *next; // Jetzt ist ListNode bekannt
};

short getLength(ListNode *head)
{
    short counter = 0;
    while (head != NULL)
    {
        counter++;
        head = head->next;

    }

    return counter;
}

int getVal(ListNode *head, short index)
{
    int val = 0;
    short counter = 0;

    ListNode *temp = head;

    while (counter != index && temp != NULL)
    {
        temp = temp->next;
        counter++;
    }

    if (temp != NULL)
    {
        return temp->val;
    }
    

    return -1;
}

void AddNode(ListNode **head, int val)
{
    ListNode *newNode = (ListNode *)malloc(sizeof(ListNode));
    newNode->val = val;
    newNode->next = NULL;

    if (*head == NULL)
    {
        *head = newNode;
    }
    else
    {
        ListNode *temp = *head;

        while (temp->next != NULL)
        {
            temp = temp->next;
        }

        temp->next = newNode;
    }
}

struct ListNode *mergeTwoLists(struct ListNode *list1, struct ListNode *list2)
{
    ListNode *resultList = NULL;

    short len1 = getLength(list1);
    short len2 = getLength(list2);

    short left_idx = 0;
    short right_idx = 0;

    while (left_idx < len1 && right_idx < len2)
    {
        int lftVal = getVal(list1, left_idx);
        int rghtVal = getVal(list2, right_idx);
        if (lftVal <= rghtVal)
        {
            AddNode(&resultList, lftVal);
            left_idx += 1;
        }
        else
        {
            AddNode(&resultList, rghtVal);
            right_idx += 1;
        }
    }

    // Check if Elements are left in the original lists and add them to the resultList
    // List1
    while (left_idx < len1)
    {
        int lftVal = getVal(list1, left_idx);
        AddNode(&resultList, lftVal);
        left_idx += 1;
    }
    // List2
    while (right_idx < len2)
    {
        int rghtVal = getVal(list2, right_idx);
        AddNode(&resultList, rghtVal);
        right_idx += 1;
    }

    return resultList;
}

int main(){
    ListNode *list1 = NULL;
    AddNode(&list1, 1);
    AddNode(&list1, 2);
    AddNode(&list1, 4);

    ListNode *list2 = NULL;
    AddNode(&list2, 1);
    AddNode(&list2, 3);
    AddNode(&list2, 4);

    ListNode *finalList = mergeTwoLists(list1, list2);

    for (short i = 0; i < 6; i++)
    {
        int val = getVal(finalList, i);
        printf("Value at %d: ", i);
        printf("%d\n", val);
    }
    
}
