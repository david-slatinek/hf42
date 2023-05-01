import React, {useEffect, useState} from 'react';

interface Book {
    isbn: string;
    title: string;
    subtitle: string;
    author: string;
    year: number;
    description: string;
    categories: string[];
    original_title: string;
    original_subtitle: string;
    original_year: number;
    translator: string;
    size: string;
    weight: string;
    pages: number;
    publisher: string;
    language: string;
    price: number;
}

export default function Display() {
    const [books, setBooks] = useState<Book[]>([]);
    useEffect(() => {
        fetch("http://localhost:8080/books", {
            method: "GET",
            mode: "cors",
            headers: {
                "Content-Type": "application/json",
            }
        })
            .then((response) => {
                console.log(response);
                return response.json()
            })
            .then((data) => {
                setBooks(data);
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, []);

    if (!books || books.length === 0) {
        return (
            <div className="container mx-auto mt-8">
                <h1 className="text-3xl font-bold mb-8 text-red-800">Found no books!</h1>
            </div>
        );
    }

    return (
        <div className="container mx-auto mt-8">
            <h1 className="text-3xl font-bold mb-8">Books</h1>
            <table className="table-auto w-full">
                <thead>
                <tr>
                    <th className="px-4 py-2">Author</th>
                    <th className="px-4 py-2">Title</th>
                    <th className="px-4 py-2">Details page</th>
                </tr>
                </thead>
                <tbody>
                {books.map((book) => (
                    <tr key={book.isbn}>
                        <td className="border px-4 py-2">{book.author}</td>
                        <td className="border px-4 py-2">{book.title}</td>
                        <td className="border px-4 py-2">
                            <a href={"http://localhost:3001/" + book.isbn} target="_blank"
                               className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                                Details
                            </a>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
}

