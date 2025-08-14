// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
将两个有序数组合并为一个有序数组
*/
contract MergeSortedArray {
    // [1,3,5,7] [2,4,6,8]
    function mergeArray(uint256[] memory arr1, uint256[] memory arr2)
        public
        pure
        returns (uint256[] memory)
    {
        uint256 len1 = arr1.length;
        uint256 len2 = arr2.length;
        uint256[] memory newArr = new uint256[](len1 + len2);
        uint256 index = 0;
        uint256 a = 0;
        uint256 b = 0;

        while (a < len1 && b < len2) {
            if (arr1[a] <= arr2[b]) {
                newArr[index] = arr1[a];
                a++;
            }else {
                newArr[index] = arr2[b];
                b++;
            }
            index++;
        }

        while (a < len1){
            newArr[index] = arr1[a];
            a++;
            index++;
        }
        while (b < len2){
            newArr[index] = arr2[b];
            b++;
            index++;
        }


        // for (uint256 i = 0; i < len1; i++) {
        //     newArr[index] = arr1[i];
        //     index++;
        // }
        // for (uint256 i = 0; i < len2; i++) {
        //     newArr[index] = arr2[i];
        //     index++;
        // }

        // uint256 len3 = newArr.length;
        // for (uint256 i = 0; i < len3 - 1; i++) {
        //     for (uint256 j = 0; j < len3 - i - 1; j++) {
        //         if (newArr[j] > newArr[j + 1]) {
        //             uint256 temp = newArr[j];
        //             newArr[j] = newArr[j + 1];
        //             newArr[j + 1] = temp;
        //         }
        //     }
        // }

        return newArr;
    }
}
