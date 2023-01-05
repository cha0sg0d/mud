// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

// External
import { DSTest } from "ds-test/test.sol";
import { World } from "solecs/World.sol";
import { IUint256Component } from "solecs/interfaces/IUint256Component.sol";
import { console } from "forge-std/console.sol";
import { LibStorage as s } from "../libraries/LibStorage.sol";
import { SampleSystem, ID as SampleSystemID } from "./SampleSystem.sol";

contract SampleTest is DSTest {
  IUint256Component public systems;
  World internal world;
  IUint256Component public components;
  SampleSystem system;

  function setUp() public {
    world = new World();
    components = world.components();
    systems = world.systems();

    // Place storage values in SampleTest's storage
    s.initSystem(world, components);
    system = new SampleSystem(world, address(components));
  }

  function testLibStorageTest() public {
    assertEq(address(world), address(s.world()));
    assertEq(address(components), address(s.comp()));
  }

  function testLibStorageSample() public {
    system.executeTyped();
  }
}
