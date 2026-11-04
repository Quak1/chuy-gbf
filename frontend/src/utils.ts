import type { Item, SlotItem, Team } from "./models";

const baseURL =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img/sp/assets";
const lowBaseURl =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img_low/sp/assets";

interface Config {
  item: Item;
  highRes: boolean;
  portrait: boolean;
  party: boolean;
  square: boolean;
}
export function getItemImageURL({
  item,
  highRes,
  portrait,
  party,
  square,
}: Config): string {
  let url = highRes ? baseURL : lowBaseURl;

  switch (item.category) {
    case "character":
      if (item.element === "any" && item.type === "") {
        // is Main Character
        return portrait
          ? `${baseURL}/leader/raid_normal/${item.image_name}.jpg`
          : `${url}/leader/m/${item.id}_01.jpg`;
      }
      return `${url}/npc/${portrait ? "f" : square ? "s" : "m"}/${item.image_name}.jpg`;
    case "summon":
      return `${url}/summon/${portrait ? "party_main" : party ? "party_sub" : square ? "s" : "m"}/${item.image_name}.jpg`;
    case "weapon":
      return `${url}/weapon/${portrait ? "ls" : square ? "s" : "m"}/${item.image_name}.jpg`;
    default:
      console.error("item has no category", item);
  }

  return "";
}

interface EmptyConfig {
  category: string;
  portrait: boolean;
}
export function getEmptyImagePath({ category, portrait }: EmptyConfig): string {
  switch (category) {
    case "character":
      return `/img/empty_character${portrait ? "_portrait" : ""}.jpg`;
    case "summon":
      return `/img/empty_summon${portrait ? "_portrait.png" : ".jpg"}`;
    case "weapon":
      return `/img/empty_weapon${portrait ? "_portrait.jpg" : ".png"}`;
    default:
      console.error("unknown category:", category);
  }

  return "";
}

export function getEmptyTeam(): Team {
  return {
    id: 0,
    name: "",
    url: "",
    character: new Array(7).fill(null),
    summon: new Array(8).fill(null),
    weapon: new Array(11).fill(null),
  };
}

export function sliceTeam(items: SlotItem[], start: number, end: number) {
  return items
    .slice(start, end)
    .map((item, i) =>
      item !== null ? item : { position: i + start, name: "-" },
    );
}
