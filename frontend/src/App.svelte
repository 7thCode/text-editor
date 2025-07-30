<script>
  import { 
    NewFile, 
    OpenFileDialog,
    SaveFile, 
    SaveFileDialog,
    GetHomeDir 
  } from '../wailsjs/go/main/App.js';

  let currentFile = {
    content: '',
    path: '',
    name: '新規ファイル'
  };
  
  let isModified = false;
  let homeDir = '';
  let textareaElement;

  // アプリケーション起動時にホームディレクトリを取得
  async function initApp() {
    try {
      homeDir = await GetHomeDir();
    } catch (error) {
      console.error('ホームディレクトリの取得に失敗しました:', error);
      homeDir = '.';
    }
  }

  // File メニュー機能
  async function handleNewFile() {
    if (isModified) {
      if (confirm('変更が保存されていません。新規ファイルを作成しますか？')) {
        currentFile = await NewFile();
        isModified = false;
      }
    } else {
      currentFile = await NewFile();
      isModified = false;
    }
  }

  async function handleOpenFile() {
    try {
      const result = await OpenFileDialog();
      currentFile = result;
      isModified = false;
    } catch (error) {
      console.error('ファイルを開けませんでした:', error);
      alert('ファイルを開けませんでした: ' + error.message);
    }
  }

  async function handleSaveFile() {
    if (!currentFile.path) {
      await handleSaveFileAs();
      return;
    }
    
    try {
      await SaveFile(currentFile.path, currentFile.content);
      isModified = false;
      alert('ファイルを保存しました');
    } catch (error) {
      console.error('ファイルの保存に失敗しました:', error);
      alert('ファイルの保存に失敗しました: ' + error.message);
    }
  }

  async function handleSaveFileAs() {
    try {
      const result = await SaveFileDialog(currentFile.content);
      currentFile = result;
      isModified = false;
      alert('ファイルを保存しました');
    } catch (error) {
      console.error('ファイルの保存に失敗しました:', error);
      alert('ファイルの保存に失敗しました: ' + error.message);
    }
  }

  // Edit メニュー機能
  function handleUndo() {
    if (textareaElement) {
      document.execCommand('undo');
    }
  }

  function handleCut() {
    if (textareaElement) {
      document.execCommand('cut');
    }
  }

  function handleCopy() {
    if (textareaElement) {
      document.execCommand('copy');
    }
  }

  function handlePaste() {
    if (textareaElement) {
      document.execCommand('paste');
    }
  }

  // テキストエリアの内容が変更されたとき
  function handleContentChange(event) {
    const target = event.target;
    currentFile.content = target.value;
    isModified = true;
  }

  // キーボードショートカット
  function handleKeydown(event) {
    if (event.ctrlKey || event.metaKey) {
      switch (event.key) {
        case 'n':
          event.preventDefault();
          handleNewFile();
          break;
        case 'o':
          event.preventDefault();
          handleOpenFile();
          break;
        case 's':
          event.preventDefault();
          if (event.shiftKey) {
            handleSaveFileAs();
          } else {
            handleSaveFile();
          }
          break;
        case 'z':
          event.preventDefault();
          handleUndo();
          break;
        case 'x':
          event.preventDefault();
          handleCut();
          break;
        case 'c':
          event.preventDefault();
          handleCopy();
          break;
        case 'v':
          event.preventDefault();
          handlePaste();
          break;
      }
    }
  }

  // コンポーネントの初期化
  initApp();
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="app-container">
  <!-- メニューバー -->
  <header class="menu-bar">
    <div class="menu-title">
      <span class="material-icons">description</span>
      <span>テキストエディター</span>
      {#if isModified}
        <span class="modified-indicator">*</span>
      {/if}
    </div>
    
    <!-- File メニュー -->
    <nav class="menu-items">
      <div class="menu-group">
        <span class="menu-label">File</span>
        <button class="menu-button" on:click={handleNewFile}>
          <span class="material-icons">add</span>
          新規作成
        </button>
        
        <button class="menu-button" on:click={handleOpenFile}>
          <span class="material-icons">folder_open</span>
          開く
        </button>
        
        <button class="menu-button" on:click={handleSaveFile}>
          <span class="material-icons">save</span>
          保存
        </button>
        
        <button class="menu-button" on:click={handleSaveFileAs}>
          <span class="material-icons">save_as</span>
          名前を付けて保存
        </button>
      </div>

      <!-- Edit メニュー -->
      <div class="menu-group">
        <span class="menu-label">Edit</span>
        <button class="menu-button" on:click={handleUndo}>
          <span class="material-icons">undo</span>
          元に戻す
        </button>
        
        <button class="menu-button" on:click={handleCut}>
          <span class="material-icons">content_cut</span>
          切り取り
        </button>
        
        <button class="menu-button" on:click={handleCopy}>
          <span class="material-icons">content_copy</span>
          コピー
        </button>
        
        <button class="menu-button" on:click={handlePaste}>
          <span class="material-icons">content_paste</span>
          貼り付け
        </button>
      </div>
    </nav>
  </header>

  <!-- ファイル情報 -->
  <div class="file-info">
    <span class="file-name">{currentFile.name}</span>
    {#if currentFile.path}
      <span class="file-path">{currentFile.path}</span>
    {/if}
  </div>

  <!-- テキストエディター -->
  <main class="editor-container">
    <textarea
      bind:this={textareaElement}
      class="text-editor"
      bind:value={currentFile.content}
      on:input={handleContentChange}
      placeholder="ここにテキストを入力してください..."
      spellcheck="false"
    ></textarea>
  </main>

  <!-- ステータスバー -->
  <footer class="status-bar">
    <span class="status-text">
      {#if isModified}
        変更済み
      {:else}
        保存済み
      {/if}
    </span>
    <span class="shortcuts">
      Ctrl+N: 新規作成 | Ctrl+O: 開く | Ctrl+S: 保存 | Ctrl+Shift+S: 名前を付けて保存 | 
      Ctrl+Z: 元に戻す | Ctrl+X: 切り取り | Ctrl+C: コピー | Ctrl+V: 貼り付け
    </span>
  </footer>
</div>

<style>
  .app-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    background-color: var(--md-sys-color-surface);
    color: var(--md-sys-color-on-surface);
  }

  .menu-bar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 16px;
    height: 64px;
    background-color: var(--md-sys-color-primary);
    color: var(--md-sys-color-on-primary);
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .menu-title {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 18px;
    font-weight: 500;
  }

  .modified-indicator {
    color: var(--md-sys-color-error);
    font-weight: bold;
  }

  .menu-items {
    display: flex;
    gap: 24px;
  }

  .menu-group {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .menu-label {
    font-size: 14px;
    font-weight: 500;
    margin-right: 8px;
    opacity: 0.9;
  }

  .menu-button {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 8px 12px;
    border: none;
    border-radius: 4px;
    background-color: transparent;
    color: var(--md-sys-color-on-primary);
    cursor: pointer;
    font-size: 14px;
    transition: background-color 0.2s;
  }

  .menu-button:hover {
    background-color: rgba(255, 255, 255, 0.1);
  }

  .menu-button:active {
    background-color: rgba(255, 255, 255, 0.2);
  }

  .file-info {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 8px 16px;
    background-color: var(--md-sys-color-surface-variant);
    color: var(--md-sys-color-on-surface-variant);
    font-size: 14px;
    border-bottom: 1px solid var(--md-sys-color-outline-variant);
  }

  .file-name {
    font-weight: 500;
  }

  .file-path {
    color: var(--md-sys-color-outline);
    font-family: monospace;
  }

  .editor-container {
    flex: 1;
    padding: 16px;
    overflow: hidden;
  }

  .text-editor {
    width: 100%;
    height: 100%;
    padding: 16px;
    border: 1px solid var(--md-sys-color-outline-variant);
    border-radius: 8px;
    background-color: var(--md-sys-color-surface);
    color: var(--md-sys-color-on-surface);
    font-family: 'Roboto Mono', 'Consolas', 'Monaco', monospace;
    font-size: 14px;
    line-height: 1.6;
    resize: none;
    outline: none;
    transition: border-color 0.2s;
  }

  .text-editor:focus {
    border-color: var(--md-sys-color-primary);
  }

  .text-editor::placeholder {
    color: var(--md-sys-color-outline);
  }

  .status-bar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 8px 16px;
    background-color: var(--md-sys-color-surface-variant);
    color: var(--md-sys-color-on-surface-variant);
    font-size: 12px;
    border-top: 1px solid var(--md-sys-color-outline-variant);
  }

  .status-text {
    font-weight: 500;
  }

  .shortcuts {
    color: var(--md-sys-color-outline);
  }

  /* レスポンシブデザイン */
  @media (max-width: 768px) {
    .menu-bar {
      flex-direction: column;
      height: auto;
      padding: 8px 16px;
    }

    .menu-items {
      flex-wrap: wrap;
      justify-content: center;
      gap: 16px;
    }

    .menu-group {
      flex-direction: column;
      align-items: flex-start;
    }

    .file-info {
      flex-direction: column;
      align-items: flex-start;
      gap: 4px;
    }

    .status-bar {
      flex-direction: column;
      gap: 4px;
      text-align: center;
    }
  }
</style>
