// Renames the section headings `changeset version` writes to the ones the rest
// of CHANGELOG.md uses. Changesets names sections after the bump level and has
// no option to change that, so this runs right after it (`pnpm run version`).

import { readFileSync, writeFileSync } from 'node:fs';

const headings = {
  '### Major Changes': '### Breaking changes',
  '### Minor Changes': '### Features',
  '### Patch Changes': '### Fixes',
};

const path = new URL('../CHANGELOG.md', import.meta.url);
const changelog = readFileSync(path, 'utf8');
const renamed = changelog
  .split('\n')
  .map((line) => headings[line] ?? line)
  .join('\n');
writeFileSync(path, renamed);
