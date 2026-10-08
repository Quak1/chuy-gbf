import type { Item } from "./models";

const baseURL =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img/sp/assets";
const lowBaseURl =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img_low/sp/assets";

export function getItemImageURL(item: Item, useHighRes = false): string {
  let url = useHighRes ? baseURL : lowBaseURl;

  switch (item.category) {
    case "character":
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
