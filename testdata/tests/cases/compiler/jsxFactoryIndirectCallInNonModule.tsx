// @target: es2015
// @module: commonjs
// @jsx: react-jsx
// @moduleDetection: legacy
// file is intentionally not a module - tests that jsx factory calls use indirect call pattern
class Component {
    render() {
        return <div>{null/* preserved */}</div>;
    }
}
