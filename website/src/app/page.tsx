import Image from "next/image";

export default function Home() {
  return (
    <div className="h-screen flex items-center justify-center">
      <main className="flex items-center justify-start flex-col gap-y-20">
        <div className="flex flex-col items-start justify-center gap-y-8">
          <h1 className="text-8xl font-bold text-white">
            config-lsp
          </h1>
          <p className="text-2xl text-gray-400">
            The Language Server for your configuration files
          </p>
        </div>
        <div className="flex items-center justify-center">
          <button className="px-6 py-3 font-bold text-white rounded-lg bg-primary hover:bg-gray-900 transition-colors cursor-pointer">
            Download for VS Code
          </button>
        </div>
      </main>
    </div>
  );
}
