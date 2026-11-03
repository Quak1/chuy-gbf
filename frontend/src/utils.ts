import type { Item, SlotItem, Team } from "./models";

const baseURL =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img/sp/assets";
const lowBaseURl =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img_low/sp/assets";

interface Config {
  item: Item;
  highRes: boolean;
  portrait: boolean;
}
export function getItemImageURL({ item, highRes, portrait }: Config): string {
  let url = highRes ? baseURL : lowBaseURl;

  switch (item.category) {
    case "character":
      if (item.element === "any" && item.type === "") {
        // is Main Character
        return portrait
          ? `${baseURL}/leader/raid_normal/${item.image_name}.jpg`
          : `${url}/leader/m/${item.id}_01.jpg`;
      }
      return `${url}/npc/${portrait ? "f" : "m"}/${item.image_name}.jpg`;
    case "summon":
      return `${url}/summon/${portrait ? "party_main" : "m"}/${item.image_name}.jpg`;
    case "weapon":
      return `${url}/weapon/${portrait ? "ls" : "m"}/${item.image_name}.jpg`;
    default:
      console.error("item has no category", item);
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
