
 export default function Header() {
    return (
        <header className="bg-gray-800 text-white p-4">
            <div className="container mx-auto flex justify-between items-center">
                <h1 className="text-2xl font-bold"><a href="/">Kira.</a></h1>
                <nav>
                    <ul className="flex space-x-4">
                        <li><a href="/about" className="hover:text-gray-400">About</a></li>
                        <li><a href="/contact" className="hover:text-gray-400">Contact</a></li>
                    </ul>
                </nav>
                <div className="flex space-x-4">
                    <a href="/login" className="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded">Login</a>
                    <a href="/signup" className="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded">Sign Up</a>
                </div>
            </div>
        </header>
    );
}