"use client";
import Image from "next/image";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-400 to-blue-500 flex flex-col">
      {/* Header */}
      <header className="flex items-center justify-between px-8 py-6">
        <div className="flex items-center gap-2">
          <Image src="/window.svg" alt="GoCloud Logo" width={32} height={32} />
          <span className="text-white font-bold text-lg">GoCloud</span>
        </div>
        <nav className="flex items-center gap-8">
          <a href="#features" className="text-white hover:underline">Features</a>
          <a href="#about" className="text-white hover:underline">About</a>
          <a href="#contact" className="text-white hover:underline">Contact</a>
          <button
            className="ml-4 px-5 py-2 border border-white rounded-lg text-white hover:bg-white hover:text-blue-500 transition"
            onClick={() => {
              window.location.href = "http://localhost:3050/auth/google";
            }}
          >
            Sign In
          </button>
        </nav>
      </header>

      {/* Main Content */}
      <main className="flex flex-1 flex-col items-center justify-center text-center px-4">
        <Image src="/window.svg" alt="Cloud Upload" width={90} height={90} className="mb-6" />
        <h1 className="text-6xl font-bold text-white mb-2">GoCloud</h1>
        <h2 className="text-2xl text-white tracking-widest mb-6">STORAGE</h2>
        <p className="text-white text-lg max-w-xl mb-10">
          Secure, reliable, and lightning-fast cloud storage solutions. Store, sync, and share your files with confidence.
        </p>
        <div className="flex gap-4 justify-center">
          <button className="bg-white text-blue-500 font-semibold px-8 py-3 rounded-lg shadow hover:bg-blue-100 transition">Get Started</button>
          <button className="bg-transparent border border-white text-white font-semibold px-8 py-3 rounded-lg hover:bg-white hover:text-blue-500 transition">Learn More</button>
        </div>
      </main>
    </div>
  );
}