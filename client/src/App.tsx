import './App.css'; 
import { Greet } from './components/Greet'

function App() {
  return (
    <div className="App">
      <Greet name='Bradley' messageCount={100} isLoggedIn={true} />
    </div>
  );
}

export default App;
