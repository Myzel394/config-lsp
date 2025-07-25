import SupportedConfig from "@/components/SupportedConfig";
import { DiTerminal } from "react-icons/di";
import { FaBitcoin, FaThList } from "react-icons/fa";
import { FaHardDrive } from "react-icons/fa6";
import { LuTerminal } from "react-icons/lu";
import { MdMail } from "react-icons/md";
import { SiWireguard } from "react-icons/si";

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
        <div className="flex flex-col items-center justify-center gap-y-4">
          <button className="px-6 py-3 font-bold text-white rounded-lg bg-primary hover:bg-gray-900 transition-colors cursor-pointer">
            Download for VS Code
          </button>
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
      </main>
    </div>
  );
}

