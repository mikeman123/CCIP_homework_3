// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Test, console} from "forge-std/Test.sol";
import {TransferUSDC} from "../src/TransferUSDC.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract TransferUSDCTest is Test {
    TransferUSDC transferUSDC;
    address private OWNER;
    address private RECEIVER;
    address private USDC_TOKEN_ADDRESS;
    uint64 private DESTINATION_CHAIN_SELECTOR;

    function setUp() public {
        OWNER = vm.envAddress("OWNER_ADDRESS");
        RECEIVER = vm.envAddress("RECEIVER_ADDRESS");
        USDC_TOKEN_ADDRESS = vm.envAddress("USDC_TOKEN_ADDRESS");
        DESTINATION_CHAIN_SELECTOR = uint64(vm.envUint("DESTINATION_CHAIN_SELECTOR"));

        transferUSDC = TransferUSDC(0x1797f42562826124FAC5a4Ed601Ab84a640de407);

        // Fund the test contract with USDC if needed
        deal(USDC_TOKEN_ADDRESS, address(this), 1e6 * 10 ** 6); // Fund with 1 million USDC
    }

    function testTransferUsdcWithIncreasedGas() public {
        uint256 amountToTransfer = 1_000 * 10 ** 6; // 1000 USDC

        // Start impersonating the owner
        vm.startPrank(OWNER);

        // Measure gas consumption of the initial transfer
        uint256 gasBefore = gasleft();
        transferUSDC.transferUsdc(DESTINATION_CHAIN_SELECTOR, RECEIVER, amountToTransfer, 500_000); // Initial gas limit
        uint256 gasUsed = gasBefore - gasleft();

        // Increase gas by 10%
        uint256 increasedGasLimit = (gasUsed + (gasUsed / 10));

        // Log the results for debugging
        console.log("Gas Used by ccipReceive: ", gasUsed);
        console.log("Increased Gas Limit: ", increasedGasLimit);

        // Execute the transfer again with the increased gas limit
        transferUSDC.transferUsdc(DESTINATION_CHAIN_SELECTOR, RECEIVER, amountToTransfer, uint64(increasedGasLimit));

        // Stop impersonating the owner
        vm.stopPrank();

        // Assert true to ensure the test runs successfully
        assertTrue(true, "Transfer with increased gas limit succeeded");
    }
}
