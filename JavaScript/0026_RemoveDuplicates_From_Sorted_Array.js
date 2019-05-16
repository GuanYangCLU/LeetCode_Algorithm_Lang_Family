/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    var arr = 1;
    for (let i=1;i<nums.length;i++) {
        if (nums[i] != nums[i-1]) nums[arr++] = nums[i];
        // console.log(arr);
    }
    return(arr);
};