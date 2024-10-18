using System;

public class Solution {
    public bool IsValid(string s){
        if (s.Length == 0)
        {
            return true;
        }
        Stack<int> stack = new Stack<int>();
        bool valid = false;
    
        for (int i = 0; i < s.Length; i++)
        {
            int lastItem = stack.Pop();
            switch (s[i])
            {
                case '(':
                    stack.Push(1);
                    break;
                case '[':
                    stack.Push(2);
                    break;
                case '{':
                    stack.Push(3);
                    break;
                case ')':
                    if (lastItem == 0)
                    {
                        valid = true;
                    }
                    else
                    {
                        return false;
                    }
                    break;
                case ']':
                    if (lastItem == 1)
                    {
                        valid = true;
                    }
                    else
                    {
                        return false;
                    }
                    break;
                case '}':
                    if (lastItem == 2)
                    {
                        valid = true;
                    }
                    else
                    {
                        return false;
                    }
                    break;
            }
        }

        return valid;
    }
}