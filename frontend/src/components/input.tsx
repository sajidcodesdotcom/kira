type inputTypes = {
    type: string;
    id: string;
    placeholder: string;
    className: string;
    value: string
    handleChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
}

export default function Input({ type, id, placeholder, className, value, handleChange }: inputTypes) {
    return (
        <input required placeholder={placeholder}
        autoComplete="on"
        type={type}
        id={id}
        value={value}
        className={`mt-1 block p-2 w-full border border-gray-300 rounded-md shadow-sm focus:ring focus:ring-blue-500 ${className}`}
        onChange={handleChange}
        />
    )
}