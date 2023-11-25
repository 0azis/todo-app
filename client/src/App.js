import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import {NavBar} from "./components/Navbar/navbar";
import {Main} from "./components/Main/main";
import {SignUp} from "./components/SignUp/signUp";
import {SignIn} from "./components/SignIn/signIn";
import {GetNotes} from "./components/Note/note";

function App() {
  return (
      <Router>
          <Routes >

                <Route path="/" element={
                        [<NavBar />, <Main />]
                }/>

              <Route path="/signup" element={
                  [<NavBar />, <SignUp />]
              } />
              <Route path="/signin" element={
                  [<NavBar />, <SignIn />]
              } />
              <Route path="/notes" element={
                  [<NavBar />, <GetNotes />]
              }/>
          </Routes>


      </Router>
  );
}

export default App;
