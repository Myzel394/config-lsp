import Image from "next/image";
import Link from "next/link";
import { FC } from "react";

const Navbar: FC<{}> = () => {
    return (
        // Aesthetic navigation bar in the center
        <nav className="flex items-center justify-center p-4 gap-x-20 text-white">
            <ul className="flex space-x-6 text-gray-300" id="nav-links">
                <li>
                    <Link
                        href="https://marketplace.visualstudio.com/items?itemName=myzel394.config-lsp"
                        className="hover:text-gray-400"
                        rel="noopener noreferrer"
                        target="_blank"
                    >Download for VS Code</Link>
                </li>
                <li className="border-l border-gray-600 h-6 pointer-events-none"></li>
                <li>
                    <Link
                        href="https://github.com/Myzel394/config-lsp"
                        className="hover:text-gray-400"
                        rel="noopener noreferrer"
                        target="_blank"
                    >GitHub</Link>
                </li>
                <li className="border-l border-gray-600 h-6 pointer-events-none"></li>
                <li>
                    <Link
                        href="https://git.myzel394.app/Myzel394/config-lsp"
                        className="hover:text-gray-400"
                        rel="noopener noreferrer"
                        target="_blank"
                    >Gitea (self hosted)</Link>
                </li>
            </ul>
        </nav>
    )
}

export default Navbar;
