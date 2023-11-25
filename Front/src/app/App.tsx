import './App.css'
import Sidebar from "../widgets/Sidebar";
import { ClerkProvider, RedirectToSignIn, SignedIn, SignedOut } from "@clerk/clerk-react";
import { dark } from "@clerk/themes";
import AppRouter from "./providers/router/ui/AppRouter.tsx";

function App() {
    if (!import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY) {
        throw new Error("Missing Publishable Key")
    }
    const clerkPubKey = import.meta.env.VITE_REACT_APP_CLERK_PUBLISHABLE_KEY;

  return (
    <>
    <ClerkProvider publishableKey={clerkPubKey} appearance={{
        baseTheme: dark,
        signIn: { baseTheme: dark },
    }}
    >
        <SignedIn>
        <div className="bg-gray-900 text-4xl flex flex-row">
                    <Sidebar isOpen={true} />
                    <div className="w-10/12 mx-auto p-3 h-screen text-white">
                        <AppRouter />
                    </div>
            </div>
        </SignedIn>
        <SignedOut>
            <RedirectToSignIn />
        </SignedOut>
    </ClerkProvider>
    </>
  )
}

export default App
