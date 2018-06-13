pragma solidity ^0.4.0;

contract Ballot {
    int x = 1;

    function test() public payable {
        x = x + 2;
        return;
    }
    function () public payable {
        x = x + 1;
        return;
    }
}
