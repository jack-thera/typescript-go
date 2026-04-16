// @noEmit: true

// @filename: test.ts
import type { Outer as Target } from "./a/outer";
import type { Outer as Source } from "./b/outer";

// Should error: b/leaf.ts Leaf is missing 'id' required by a/leaf.ts Leaf.
// tsc reports this. tsgo misses it when types are nested 3+ import levels deep with arrays.
declare function getSource(): Source;
const target: Target = getSource();

// @filename: a/outer.ts
import type { Inner } from "./inner";

export interface Outer {
  inners: Inner[];
}

// @filename: a/inner.ts
import type { Mid } from "./leaf";

export interface Inner {
  mids: Mid[];
}

// @filename: a/leaf.ts
export interface Leaf {
  id: string;
}

export interface Mid {
  leaves: Leaf[];
}

// @filename: b/outer.ts
import type { Inner } from "./inner";

export interface Outer {
  inners: Inner[];
}

// @filename: b/inner.ts
import type { Mid } from "./leaf";

export interface Inner {
  mids: Mid[];
}

// @filename: b/leaf.ts
export interface Leaf {
}

export interface Mid {
  leaves: Leaf[];
}
