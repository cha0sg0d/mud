// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

// External
import "solecs/System.sol";
import { IWorld } from "solecs/interfaces/IWorld.sol";
import { LibStorage } from "../libraries/LibStorage.sol";

abstract contract StdSystem is System {
  constructor(IWorld _world, address _components) System(_world, _components) {
    LibStorage.initSystem(_world, IUint256Component(_components));
  }
}
