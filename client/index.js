import './src/App.scss';
import App from './src/App.svelte';

window.app = new App({
  target: document.getElementsByTagName('app')[0]
});
