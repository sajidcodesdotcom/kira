import Input from "../components/input";


export default function SignUpPage() {
    return (
        <div className="flex flex-col items-center justify-center overflow-hidden h-screen bg-gray-100">
        <h1 className="text-3xl font-bold mb-4">Sign Up</h1>
        <form className="bg-white p-6 rounded shadow-md w-96">
            <div className="mb-4">
                <Input type={"text"} id={"full-name"} placeholder={"Full name"} className={""} />
            </div>
            <div className="mb-4">
                <Input type={"text"} id={"username"} placeholder={"Username"} className={""} />
            </div>
            <div className="mb-4">
                <Input type={"email"} id={"email"} placeholder={"Email"} className={""} />
            </div>
            <div className="mb-4">
                <Input type={"password"} id={"password"} placeholder={"Password"} className={""} />
            </div>
            <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded p-1 hover:bg-blue-700">Sign Up</button>
        </form>
        </div>
    );
    }