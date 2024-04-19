#include <stdio.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int* twoSum(int* nums, int numsSize, int target) {
    int* result = (int*)malloc(2 * sizeof(int));

    if (!result)
    {
        printf("Malloc has failed!");
        int arrayFail[2] = {0,0};
        return arrayFail;
    }

    for(int i = 0; i< numsSize; i++){
        for (int j = i+1; j < numsSize; j++)
        {
            if (nums[i] + nums[j] == target)
            {
                result[0] = i;
                result[1] = j;

                return result;
            }   
        }   
    }

    return result;
}

int* two2Sum(int* nums, int numsSize, int target) {
    int* result = (int*)malloc(2 * sizeof(int));

    if (!result)
    {
        printf("Malloc has failed!");
        result[0] = 0;
        result[1] = 0;
        return result;
    }

    for(int i = 0; i< numsSize; i++){
        if(i+1 < numsSize){
            if(nums[i] + nums[i+1] == target){
                result[0] = i;
                result[1] = i+1;
                return result;
            }
        }   
    }
    return result;
}

int main(){
    int nums[4] = {2,7,11,15};
    int target = 9;

    int* result = two2Sum(nums, 4, target);

    printf("Index is: %d, %d", result[0], result[1]);
}