//// [tests/cases/compiler/jsxFactoryIndirectCallInNonModule.tsx] ////

//// [jsxFactoryIndirectCallInNonModule.tsx]
// file is intentionally not a module - tests that jsx factory calls use indirect call pattern
class Component {
    render() {
        return <div>{null/* preserved */}</div>;
    }
}


//// [jsxFactoryIndirectCallInNonModule.js]
"use strict";
// file is intentionally not a module - tests that jsx factory calls use indirect call pattern
class Component {
    render() {
        return (0, _a.jsx)("div", { children: null /* preserved */ });
    }
}
