// Writes each changeset's summary as a CHANGELOG.md bullet. Changesets' default
// prefixes every bullet with the commit hash, which the hand-written entries in
// CHANGELOG.md never had. The section headings are renamed afterwards by
// rename-changelog-headings.mjs.

async function getReleaseLine(changeset) {
  const [firstLine, ...rest] = changeset.summary.split('\n').map((line) => line.trimEnd());
  const body = rest.map((line) => (line ? `  ${line}` : line)).join('\n');
  return `\n\n- ${firstLine}${body ? `\n${body}` : ''}`;
}

// The provider is the only package, so no release ever bumps a dependency.
async function getDependencyReleaseLine() {
  return '';
}

export default { getReleaseLine, getDependencyReleaseLine };
