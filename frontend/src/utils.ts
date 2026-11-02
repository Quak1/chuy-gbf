import type { Item, SlotItem, Team } from "./models";

const baseURL =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img/sp/assets";
const lowBaseURl =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img_low/sp/assets";

export function getItemImageURL(item: Item, useHighRes = false): string {
  let url = useHighRes ? baseURL : lowBaseURl;

  switch (item.category) {
    case "character":
      if (item.type === "") return `${url}/leader/m/${item.id}_01.jpg`;
      return `${url}/npc/m/${item.id}_01.jpg`;
    case "summon":
      return `${url}/summon/m/${item.id}.jpg`;
    case "weapon":
      return `${url}/weapon/m/${item.id}.jpg`;
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
    characters: new Array(7).fill(null),
    summons: new Array(8).fill(null),
    weapons: new Array(11).fill(null),
  };
}

export function sliceTeam(items: SlotItem[], start: number, end: number) {
  return items
    .slice(start, end)
    .map((item, i) =>
      item !== null ? item : { position: i + start, name: "-" },
    );
}
