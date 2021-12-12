import React, { useEffect } from 'react'
import "./SigninPage.css"

export default function SigninPage() {
    useEffect(() => {
        document.title = "Twidder - Sign in"
    }, [])
 
    return (
        <div>
            <h1>Sign in!</h1>
        </div>
    )
}
