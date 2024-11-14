// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

contract counter_event_emitter {
    int private count = 0;
    event Count(int count);

    function increment() public {
        count += 1;
        emit Count(count);
    }

    function getCount() public view returns (int) {
        return count;
    }
}
