// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import { IUint256Component } from "solecs/interfaces/IUint256Component.sol";
import { IWorld } from "solecs/interfaces/IWorld.sol";
import { getSystemAddressById, getAddressById } from "solecs/utils.sol";

uint256 constant worldID = uint256(keccak256("mud.world"));
uint256 constant componentsID = uint256(keccak256("mud.components"));

/**
 * This library will only work if you use `std-contracts/StdSystem.sol` for every system in your contracts
 */

library LibStorage {
  struct Layout {
    IUint256Component components;
    IWorld world;
  }

  bytes32 internal constant STORAGE_SLOT = keccak256("solecs.contracts.storage.System");

  function layout() internal pure returns (Layout storage l) {
    bytes32 slot = STORAGE_SLOT;
    assembly {
      l.slot := slot
    }
  }

  function initSystem(IWorld _world, IUint256Component _components) internal {
    Layout storage l = layout();
    l.world = _world;
    l.components = _components;
  }

  function comp() public view returns (IUint256Component) {
    return layout().components;
  }

  function world() public view returns (IWorld) {
    return layout().world;
  }
}
