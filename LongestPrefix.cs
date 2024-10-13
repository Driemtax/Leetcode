using System;
using System.Linq;

public class Solution
{
    public string LongestCommonPrefix(string[] strs)
    {
        string prefix = "";

        if (strs == null || strs.Length == 0)
        {
            return prefix;
        }

        string shortest = strs.OrderBy(s => s.Length).First();

        for (int i = 0; i < shortest.Length; i++)
        {
            foreach (string str in strs)
            {
                if (str[i] != shortest[i])
                {
                    return prefix;
                }
            }

            prefix += shortest[i];
        }

        return prefix;
    }

    public static void Main()
    {
        string[] strs = { "flower", "flow", "flight" };

        string commonPrefix = LongestCommonPrefix(strs);

        Console.WriteLine("Longest Prefix: " + commonPrefix);
    }
}