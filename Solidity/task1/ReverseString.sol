// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
反转一个字符串。输入 "abcde"，输出 "edcba"
*/
contract ReverseString {
    function reverseStr(string memory str) public pure returns (string memory) {
        //将字符串转为字节数组
        bytes memory arr = bytes(str);
        //获取长度
        uint256 len = arr.length;
        //创建一个新字节数组，用于接收反转数据
        bytes memory reversed = new bytes(len);
        for (uint256 i = 0; i < len; i++) {
            reversed[i] = arr[len - 1 - i];
        }
        //返回反转字符
        return string(reversed);
    }
}
