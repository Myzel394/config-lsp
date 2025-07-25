import Navbar from "@/components/Navbar";
import SupportedConfig from "@/components/SupportedConfig";
import Image from "next/image";
import Link from "next/link";
import { DiTerminal } from "react-icons/di";
import { FaBitcoin, FaThList } from "react-icons/fa";
import { FaHardDrive } from "react-icons/fa6";
import { LuTerminal } from "react-icons/lu";
import { MdMail } from "react-icons/md";
import { SiWireguard } from "react-icons/si";

export default function Home() {
  return (
    <>
      <Navbar />
      <div className="h-screen flex items-center justify-center gap-x-20">

        <main className="flex items-center justify-start flex-col gap-y-20">
          <div className="flex flex-col items-start justify-center gap-y-8">

            <div className="flex items-center justify-center gap-x-2">
              <Image
                src="/assets/icon.png"
                alt="logo"
                width={70}
                height={70}
                // Optical center
                className="inline-block mt-[15px]"
              />
              <h1 className="text-8xl font-bold text-white">
                config-lsp
              </h1>
            </div>

            <p className="text-2xl text-gray-400">
              The Language Server for your configuration files
            </p>
          </div>
          <div className="flex flex-col items-center justify-center gap-y-8">
            <div className="flex items-center justify-center gap-x-2">
              <button className="px-6 py-3 font-bold text-white rounded-lg bg-primary hover:bg-gray-900 transition-colors cursor-pointer">
                Download for VS Code
              </button>
              <Link
                href="https://github.com/Myzel394/config-lsp"
                className="px-6 py-3 font-bold text-white rounded-lg cursor-pointer border-gray-800 border"
                rel="noopener noreferrer"
                target="_blank"
              >
                Check on GitHub
              </Link>
            </div>
            <div className="flex flex-col items-start justify-center gap-y-4">
              <p className="text-gray-500 text-lg">
                Supported languages:
              </p>
              <div className="grid grid-cols-3 gap-4">
                <SupportedConfig Icon={SiWireguard} name="WireGuard" />
                <SupportedConfig Icon={FaBitcoin} name="Bitcoin" />
                <SupportedConfig Icon={FaHardDrive} name="fstab" />
                <SupportedConfig Icon={DiTerminal} name="SSH" />
                <SupportedConfig Icon={LuTerminal} name="SSH Daemon" />
                <SupportedConfig Icon={MdMail} name="aliases" />
                <SupportedConfig Icon={FaThList} name="hosts" />
              </div>
            </div>
          </div>
        </main>

        <div>
          <video width="640" height="100%" autoPlay loop muted preload="auto" className="rounded-[45px] shadow-lg" id="preview-video">
            <source
              src="/assets/preview-video.mp4"
              type="video/mp4"
            />
            Your browser does not support the video tag.
          </video>
        </div>

      </div>
    </>
  );
}

