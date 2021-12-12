import './App.css';
import React from 'react';
import Home from './pages/Home/Home';
import SigninPage from "./pages/SigninPage/SigninPage"
import SignupPage from './pages/SignupPage/SignupPage';
import NotFound from './pages/NotFound/NotFound'
import { BrowserRouter, Routes, Route } from "react-router-dom"

function App() {
  return (
      <BrowserRouter>
          <Routes>
             <Route exact path="/" element={<Home/>}/>
               {/* <Route exact path="/set-reminder" component={SetReminderPage}/> */}
              <Route exact path="/signin" element={<SigninPage/>}/>
              <Route exact path="/signup" element={<SignupPage/>}/>
              <Route path="*" element={<NotFound/>}/>
          </Routes>
      </BrowserRouter>
  );
}

export default App;
