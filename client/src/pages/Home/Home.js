import React, { useEffect } from 'react'
import "./Home.css"
import bird from "../../img/bird2.png"
import GoogleLogin from 'react-google-login';
import GoogleButton from '../../components/GoogleButton/GoogleButton';

export default function Home() {
    useEffect(()=>{
        document.title = "Twidder. It's what's happening"
    },[])

    const responseGoogle = (response) => {
        console.log(response);
    }

    return (
        <div className="home">
            <div className="signup">
                <div className="form">
                    <img src={bird} alt="bird" height="50px" width="50px" id="topBird" className="bird-img"/>
                    <br/>
                    <div className="happening-now">Happening Now</div>
                    <div className="join">Join Twidder today.</div>
                    <GoogleLogin
                        clientId="658977310896"
                        render={() => (
                            <GoogleButton/>
                        )}
                        onSuccess={responseGoogle}
                        onFailure={responseGoogle}
                        cookiePolicy={'single_host_origin'}
                    />
                    <button id="sign-up">Sign up with phone or email</button>
                    <br/>
                    <div className="already">Already have an account?</div>
                    <button onClick={null} id="sign-in">Sign in</button>
                </div>
                <div className="bird">
                    <img src={bird} alt="bird" id="bottomBird" className="bird-img"/>
                </div>
            </div>
            <footer className="footer">                
                <a href="https://github.com/darienmiller88/Better-Twitter-Clone" className="code-link">
                    View code!
                </a>
                <br/>
                © 2021 Twidder, Inc.
            </footer>
        </div>
        
    )
}
