type inputTypes = {
    type: string;
    id: string;
    placeholder: string;
    className: string;
}

export default function Input({ type, id, placeholder, className }: inputTypes) {
    return (
        <input required placeholder={placeholder}
        type={type}
        id={id}
        className={`mt-1 block p-2 w-full border border-gray-300 rounded-md shadow-sm focus:ring focus:ring-blue-500 ${className}`}
        />
    )
}