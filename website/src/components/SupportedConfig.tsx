import { FC, JSX } from "react";
import { IconType } from "react-icons";

export interface SupportedConfigProps {
    Icon: IconType;
    name: string;
}

const SupportedConfig: FC<SupportedConfigProps> = ({Icon, name}) => {
    return (
        <div className="flex items-center gap-x-2 text-gray-500">
            <Icon />
            <span className="text-lg font-semibold">
                {name}
            </span>
        </div>
    )
}

export default SupportedConfig;
