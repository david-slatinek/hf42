import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

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

export default function Detail() {
    const {isbn} = useParams();

    if (!isbn) {
        return (
            <div className="container mx-auto mt-8">
                <h1 className="text-3xl font-bold mb-8 text-red-800">Invalid isbn!</h1>
            </div>
        );
    }

    const [book, setBook] = useState<Book>();
    useEffect(() => {
        fetch("http://localhost:8080/book/" + isbn, {
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
                setBook(data);
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, []);

    if (book?.isbn == null) {
        return (
            <div className="container mx-auto mt-8">
                <h1 className="text-3xl font-bold mb-8 text-red-800">Book not found</h1>
            </div>
        );
    }

    return (
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
            <div className="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between">
                <h3 className="text-lg leading-6 font-medium text-gray-900">{book.title}</h3>
                <p className="mt-3 text-sm text-gray-500 sm:mt-0 sm:ml-6">{book.author}</p>
            </div>

            <div className="mt-8">
                <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">subtitle</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.subtitle}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Categories</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.categories.join(", ")}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">ISBN</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.isbn}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Year</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.year}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Original Year</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.original_year}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Original Title</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.original_title}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Original Subtitle</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.original_subtitle}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Translator</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.translator}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Size</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.size}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Weight</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.weight}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Pages</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.pages}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Publisher</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.publisher}</dd>
                    </div>

                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Language</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.language}</dd>
                    </div>
                    <div className="sm:col-span-1">
                        <dt className="text-sm font-medium text-gray-500">Price</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.price} EUR</dd>
                    </div>

                    <div className="sm:col-span-2">
                        <dt className="text-sm font-medium text-gray-500">Description</dt>
                        <dd className="mt-1 text-sm text-gray-900">{book.description}</dd>
                    </div>
                </dl>
            </div>
        </div>
    );
}