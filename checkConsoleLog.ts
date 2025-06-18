import fs  from 'fs';
import path  from 'path';
import readline  from 'readline';

const EXCLUDE_DIRS = ['node_modules', 'dist', 'build', 'coverage', 'packages', 'assets'];
const ROOT_DIR = '.';

const searchConsoleLogs = async (filePath: string): Promise<void> => {
  const rl = readline.createInterface({
    input: fs.createReadStream(filePath),
    crlfDelay: Infinity
  });
  let lineNumber = 0;

  for await (const line of rl) {
    lineNumber++;
    if (line.includes('console.log')) {
      console.log(`console.log found: ${filePath}:${lineNumber}`);
    }
  }
};

const fullScanDirectory = (dir: string): void => {
  try {
    const entries = fs.readdirSync(dir, { withFileTypes: true });
    for (const entry of entries) {
      const fullPath = path.join(dir, entry.name);
  
      if (entry.isDirectory()) {
        if (!EXCLUDE_DIRS.includes(entry.name)) fullScanDirectory(fullPath);
      } else {
        if (entry.name.endsWith('.js') || entry.name.endsWith('.ts')) {
          try {
            searchConsoleLogs(fullPath)
          } catch (error) {
            console.error(`Error reading file ${fullPath}:`, error)
          }
        }
      }
    }
  } catch (error) {
    console.error(`Error accessing directory ${dir}:`, error);
    return;
  }
};

fullScanDirectory(ROOT_DIR);
