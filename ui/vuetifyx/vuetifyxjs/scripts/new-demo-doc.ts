import { fileURLToPath } from 'url';
import path from 'path';
import fs from 'fs';
import inquirer from 'inquirer';
import { exec } from 'child_process';

// Get the current file path
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const rootDir = path.resolve(__dirname, '../');
const docsDir = path.join(rootDir, 'docs');
const componentsDir = path.join(docsDir, 'Components')
const sidebarFile = path.join(docsDir, 'sidebar.ts');

// Read and parse the sidebar.ts file
function readSidebar(): any[] {
  const sidebarContent = fs.readFileSync(sidebarFile, 'utf8');
  const cleanedContent = sidebarContent.replace(/export\s+default\s+/, ''); // Strip the export default line
  const sidebarObject = new Function(`return ${cleanedContent}`)(); // Use Function to safely parse the content
  return sidebarObject['/']; // Return the categories
}

// Write to the sidebar.ts file
function writeSidebar(content: string): void {
  fs.writeFileSync(sidebarFile, content, 'utf8');
}

// Check if a file already exists
function checkFileExists(filePath: string): boolean {
  return fs.existsSync(filePath);
}

// Generate a markdown template
function generateMarkdownTemplate(componentName: string): string {
  return `# ${componentName}

这是一个基本示例，你可以用 \`markdown\` 语法 和 \`vue3\`、\`vuetify\` 在此处写任何组件代码


## vuetifyx 组件开发流程

### 1.新建组件
在 \`qor5/x/vuetifyx/src/lib\` 里新建任何 vue 组件, 比如 \`qor5/x/vuetifyx/src/lib/richEditor/index.vue\`

### 2.注册组件
在 \`qor5/x/vuetifyx/src/lib/plugins/index.vue\` 里注册组件，这样注册可以确保生产环境和本地环境都可用

1. vuetify 组件不用注册，直接用就行，比如 \`v-btn\`
2. vuetifyx 组件请以 \`vx-\` 开头

\`\`\`ts
// qor5/x/vuetifyx/src/lib/plugins/index.vue

import TextField from '@/lib/Form/TextFiled.vue'

const vuetifyx = {
  install: (app: App) => {
    app.component('vx-datepicker', Datepicker)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-filter', Filter)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-textdatepicker', TextDatepicker)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-messagelistener', MessageListener)
    app.component('vx-overlay', Overlay)
    app.component('vx-text-field', TextField)
    // 在此注册你的新组件
    // app.component('vx-rich-editor', TextField)
  }
}
\`\`\`

### 3. 直接在当前文档使用

可以在当前 markdown 里倒入任何本地局部示例组件，比如

\`import VueJsonPretty from 'vue-json-pretty'\`

:::demo

\`\`\`vue
<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty'

</script>
<template>
  <v-btn color="primary">hello world</v-btn>

  <VueJsonPretty :data="value"></VueJsonPretty>
  你可以任意更改这里的代码
</template>
\`\`\`

<style scoped></style>
:::

### 4. 为组件撰写必要说明和参数
目前先随意，后期会有规范
`
}

// Add the component to the correct directory in sidebar.ts
function addComponentToSidebar(componentName: string, directory: string): void {
  const sidebarContent = readSidebar()

  const linkPath = `/Components/${componentName}/` // Adjusting for flat file structure
  const newItem = { text: componentName, link: linkPath }

  // Find the correct directory in the sidebar
  const targetCategory = sidebarContent.find((category: any) => category.text === directory)

  if (targetCategory) {
    // Insert into the existing category's items
    targetCategory.items.push(newItem)
  } else {
    // Create a new category and insert the item
    sidebarContent.push({
      text: directory,
      items: [newItem]
    })
  }

  // Update sidebar.ts
  const updatedSidebar = `export default ${JSON.stringify({ '/': sidebarContent }, null, 2)};`
  writeSidebar(updatedSidebar)

  console.log(
    `Component ${componentName} has been successfully registered in sidebar.ts at path ${linkPath}`
  )
}

// Try to open the file in VSCode
function openFileInVSCode(filePath: string) {
  exec('command -v code', (error, stdout) => {
    if (!error && stdout.trim()) {
      exec(`code ${filePath}`, (err) => {
        if (err) {
          console.log(`Failed to open the file: ${err.message}`);
        } else {
          console.log('VSCode opened the file successfully.');
        }
      });
    } else {
      console.log('VSCode "code" command not found, skipping opening the file.');
    }
  });
}

// Main function
async function main(): Promise<void> {
  const sidebarCategories = readSidebar();

  // Ask if the user wants to create a new component in an existing category or a new one
  const { useExistingCategory } = await inquirer.prompt<{ useExistingCategory: boolean }>([
    {
      type: 'confirm',
      name: 'useExistingCategory',
      message: 'Do you want to create the component in an existing category?',
      default: true,
    }
  ]);

  let finalCategory = 'demo'; // Default to demo if not using existing categories

  if (useExistingCategory) {
    // Let the user choose an existing category
    const { selectedCategory } = await inquirer.prompt<{ selectedCategory: string }>([
      {
        type: 'list',
        name: 'selectedCategory',
        message: 'Please select a category:',
        choices: sidebarCategories.map((category: any) => category.text),
      }
    ]);

    finalCategory = selectedCategory;
  } else {
    // Allow the user to create a new category
    const { newCategory } = await inquirer.prompt<{ newCategory: string }>([
      {
        type: 'input',
        name: 'newCategory',
        message: 'Enter the name for the new category:',
        validate: (input) => input ? true : 'Category name cannot be empty',
      }
    ]);

    finalCategory = newCategory;

    // Check if the category already exists in the sidebar
    const categoryExists = sidebarCategories.some((category: any) => category.text === newCategory);

    if (!categoryExists) {
      console.log(`Creating new category: ${newCategory}`);
    }
  }

  // Input the component filename and check for conflicts
  let validFileName = false;
  let componentName = '';

  while (!validFileName) {
    const { newComponentName } = await inquirer.prompt<{ newComponentName: string }>([
      {
        type: 'input',
        name: 'newComponentName',
        message: 'Enter the component documentation file name:',
        validate: (input) => input ? true : 'File name cannot be empty',
      }
    ]);

    const targetFilePath = path.join(componentsDir, newComponentName, 'index.md'); // Flat structure

    if (checkFileExists(targetFilePath)) {
      console.log(`File already exists: ${targetFilePath}, please enter a different name.`);
    } else {
      componentName = newComponentName;
      validFileName = true;

      // If the directory doesn't exist, create it
      if (!fs.existsSync(path.dirname(targetFilePath))) {
        fs.mkdirSync(path.dirname(targetFilePath), { recursive: true });
      }

      // Create the markdown file
      fs.writeFileSync(targetFilePath, generateMarkdownTemplate(componentName));
      console.log(`Successfully created new Markdown file: ${targetFilePath}`);

      // Open the file in VSCode (if available)
      openFileInVSCode(targetFilePath);

      // Register the component in the sidebar.ts
      addComponentToSidebar(componentName, finalCategory);
    }
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
