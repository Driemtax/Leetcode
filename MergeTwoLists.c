#include <stdio.h>

typedef struct
{
    int val;
    ListNode *next;
} ListNode;

short getLength(ListNode *head)
{
    short counter = 0;
    while (head->next != NULL)
    {
        counter++;
    }

    return counter;
}

int getVal(ListNode *head, short index)
{
    int val = 0;
    short counter = 0;

    ListNode *temp;

    while (counter != index)
    {
        temp = head->next;
    }

    val = temp->val;

    return val;
}

void AddNode(ListNode **head, int val)
{
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
            resultList->val = lftVal;
            left_idx += 1;
        }
        else
        {
            resultList->val = rghtVal;
            right_idx += 1;
        }
    }
}
