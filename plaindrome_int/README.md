Palindrome Int

Given an integer x, return true if x is a 
palindrome
, and false otherwise.

 

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 

Constraints:

-231 <= x <= 231 - 1
 

Follow up: Could you solve it without converting the integer to a string?


Explanation:

Here's a step-by-step explanation of the code:

Initialization:

go
Copy code
var b int
var t int = x
b: This variable will be used to store the reversed version of the input integer.
t: This variable is assigned the initial value of the input integer x. It's used later for comparison.
Check for Negative Number:

go
Copy code
if x < 0 {
    return false
}
If the input integer x is negative, the function immediately returns false since negative numbers cannot be palindromes.
Reverse the Integer:

go
Copy code
for x > 0 {
    b = b * 10 + x % 10
    x = x / 10
}
The loop iterates as long as x is greater than 0.
In each iteration, the last digit of x is obtained using the modulus operator (x % 10), and it is added to the reversed number b after shifting its digits to the left by one position (b * 10).
The last digit is removed from x by dividing it by 10 (x / 10).
This process continues until all digits of the original number x have been processed.
Check for Palindrome:


if t == b {
    return true
} else {
    return false
}
After the loop, the reversed number b is compared with the original number t.
If they are equal, the function returns true, indicating that the original number is a palindrome. Otherwise, it returns false.
The logic behind the code is to reverse the digits of the input integer and check if the reversed version is the same as the original integer. If they match, the integer is a palindrome. The code handles the special case of negative numbers, considering them not palindromic