import {LitElement, html} from 'lit';
import { myStyles as myStyles } from './styles.js';

class HelloWorld extends LitElement {
  static styles = [myStyles]

  render() {
    return html`
    <h1 class="title">👋 Hello World 🌍</h1>
    `;
  }
}

customElements.define('hello-world', HelloWorld);

class ServedBySlingshot extends LitElement {
  static styles = [myStyles]

  render() {
    return html`
    <h2 class="subtitle">Served with 💜 by Slingshot</h2>
    `;
  }
}

customElements.define('served-by-slingshot', ServedBySlingshot);

class LittleMessage extends LitElement {
  static styles = [myStyles]

  render() {
    return html`
    <h2 class="subtitle"> {{message}} </h2>
    `;
  }
}

customElements.define('little-message', LittleMessage);

class MainApp extends LitElement {
  static styles = [myStyles]

  render() {
    return html`
    <section class="container">
      <div>
        <hello-world></hello-world>
        <served-by-slingshot></served-by-slingshot>
        <little-message></little-message>
      </div>
    </section>    
    `;
  }
}

customElements.define('main-app', MainApp);